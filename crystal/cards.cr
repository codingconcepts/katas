MASTER_DECK = Deck.master

enum Suit
  Spades
  Hearts
  Diamonds
  Clubs
end

enum Num
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
end

struct Card
	property suit, num

	def initialize(@suit : Suit, @num : Num)
	end
end

struct Deck
	property dealIndex, cards

	def self.master()
		d = Deck.new
		d.cards
	end

	def self.copy()
		d = Deck.new
		d.cards = 
		d
	end

	def deal(count : Int32)
		cards.sample(count)
	end
end

d = Deck.copy