#[macro_use]
extern crate lazy_static;
extern crate chan;
extern crate rand;
extern crate num_cpus;

use std::thread;
use std::time::{Instant};
use rand::Rng;

lazy_static! {
    static ref SUITS: Vec<Suit> = vec![
        Suit::Clubs,
        Suit::Diamonds,
        Suit::Hearts,
        Suit::Spades
    ];
    
    static ref NUMBERS: Vec<Number> = vec![
        Number::Ace,
        Number::Two,
        Number::Three,
        Number::Four,
        Number::Five,
        Number::Six,
        Number::Seven,
        Number::Eight,
        Number::Nine,
        Number::Ten,
        Number::Jack,
        Number::Queen,
        Number::King
    ];

    static ref MASTER_DECK: Deck = Deck::master();
}

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
enum Suit {
	Spades,
	Hearts,
	Diamonds,
	Clubs,
}

#[derive(Debug, Copy, Clone, PartialEq, Eq, Hash)]
enum Number {
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

#[derive(Debug, Copy, Clone)]
struct Card {
	number: Number,
	suit: Suit,
}

#[derive(Clone)]
struct Deck {
	cards: Vec<Card>,
    deal_index: usize,
}

fn main() {
    let now = Instant::now();

    perf_test(num_cpus::get(), 100000000, 25);

    println!("Took {:#?}", now.elapsed());
}

fn perf_test(workers: usize, iterations: usize, cards_to_deal: usize) {
	let wg = chan::WaitGroup::new();
    for _ in 0..workers {
        wg.add(1);
        let wg = wg.clone();
        thread::spawn(move || {
            dealer(iterations / workers, cards_to_deal);
            wg.done();
        });
    }
    wg.wait();
}

fn dealer(iterations: usize, cards_to_deal: usize) {
    let mut xor = XorShift::new();

    for _ in 0..iterations {
        let mut d = Deck::new();
        for _ in 0..cards_to_deal {
            d.deal(&mut xor);
        }
    }
}

impl Deck {
    fn master() -> Deck {
        let mut d = Deck{cards: Vec::new(), deal_index: 0};
        for s in SUITS.iter() {
            for n in NUMBERS.iter() {
                d.cards.push(Card{number: *n, suit: *s});
            }
        }

        d
    }

    fn new() -> Deck {
        let d = MASTER_DECK.clone();
        d
    }

    fn deal(&mut self, xor: &mut XorShift) -> Card {
        let next = xor.between(self.deal_index, 52);
        let card = self.cards[next];

        self.cards.swap(next, self.deal_index);
        self.deal_index += 1;

        card
    }
}

struct XorShift {
	x: usize,
	y: usize,
	z: usize,
	w: usize,
}

impl XorShift {
	fn new() -> XorShift {
        let mut rng = rand::thread_rng();

        XorShift {
            x: rng.gen::<usize>(),
            y: rng.gen::<usize>(),
            z: rng.gen::<usize>(),
            w: rng.gen::<usize>(),
        }
	}

	fn between(&mut self, min: usize, max: usize) -> usize {
		self.next() % (max - min) + min
	}

	fn next(&mut self) -> usize {
		let t = self.x ^ (self.x << 11);

		self.x = self.y;
		self.y = self.z;
		self.z = self.w;

		self.w = (self.w ^ (self.w >> 19)) ^ (t ^ (t >> 8));
		self.w
	}
}

#[cfg(test)]
mod test;