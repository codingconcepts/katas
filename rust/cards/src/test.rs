use super::*;

use std::collections::HashMap;

#[test]
fn test_fastxor_between() {
	let mut xor = XorShift::new();
	
	for _ in 0..1000 {
		let b = xor.between(1, 10);
		assert!(b > 0 && b < 10);
	}
}

#[test]
fn test_master_deck_init() {
	assert_eq!(0, MASTER_DECK.deal_index);
	assert_eq!(52, MASTER_DECK.cards.len());

	let mut i = 0;
	for s in SUITS.iter() {
		for n in NUMBERS.iter() {
			let card = MASTER_DECK.cards[i];
			assert_eq!(*s, card.suit);
			assert_eq!(*n, card.number);
			i += 1;
		}
	}
}

#[test]
fn test_deck_new() {
	let clone = Deck::new();
	assert_eq!(MASTER_DECK.deal_index, clone.deal_index);

	for (i, c) in MASTER_DECK.cards.iter().enumerate() {
		assert_eq!(c.suit, clone.cards[i].suit);
		assert_eq!(c.number, clone.cards[i].number);
	}
}

#[test]
fn test_deck_deal() {
	let mut d = Deck::new();
	let mut xor = XorShift::new();

	let mut suits: HashMap<Suit, usize> = HashMap::new();
	let mut numbers: HashMap<Number, usize> = HashMap::new();
	
	for _ in 0..52 {
		let c = d.deal(&mut xor);
		
		let s = suits.entry(c.suit).or_insert(0);
		*s += 1;

		let n = numbers.entry(c.number).or_insert(0);
		*n += 1;
	}

	assert_eq!(52, d.deal_index);

	// Assert there are 13 of each suit (1 for each number).
	assert_eq!(4, suits.len());
	assert!(suits.iter().all(|(_, v)| *v == 13));

	// Assert there are 4 of each number (1 for each suit).
	assert_eq!(13, numbers.len());
	assert!(numbers.iter().all(|(_, v)| *v == 4));
}