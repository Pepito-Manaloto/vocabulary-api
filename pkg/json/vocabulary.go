package json

type Vocabulary struct {
    EnglishWord string `json:"english_word"`
    ForeignWord string `json:"foreign_word"`
}

type Languages struct {
    Hokkien *[]Vocabulary `json:"Hokkien"`
    Japanese *[]Vocabulary `json:"Japanese"`
    Mandarin *[]Vocabulary `json:"Mandarin"`
}

// TODO: refactor hard-coded languages
type Vocabularies struct {
    Languages *Languages `json:"languages"`
    RecentlyAddedCount string `json:"recently_added_count"`
}
