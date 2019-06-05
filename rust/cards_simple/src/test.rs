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

	for i in 0..52 {
		assert_eq!(i as u8, MASTER_DECK.cards[i]);
	}
}

#[test]
fn test_deck_new() {
	let clone = Deck::new();
	assert_eq!(MASTER_DECK.deal_index, clone.deal_index);

	for i in 0..52 {
		assert_eq!(i as u8, clone.cards[i]);
	}
}

#[test]
fn test_deck_deal() {
	let mut d = Deck::new();
	let mut xor = XorShift::new();

	let mut dealt: HashMap<u8, ()> = HashMap::new();
	
	for _ in 0..52 {
		let c = d.deal(&mut xor);
		
		if dealt.contains_key(&c) {
			panic!("dealt {} twice", c);
		}
		dealt.insert(c, ());
	}

	assert_eq!(52, d.deal_index);
	assert_eq!(52, dealt.len());
}