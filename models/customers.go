package models

type VehiculeSimple struct {
	Libelle_court_gamme *string `bson:"libelle_court_gamme,omitempty" json:"libelle_court_gamme,omitempty"`
}

type Customer struct {
	IDClient       uint32           `bson:"id_client" json:"id_client"`
	Nom            string           `bson:"nom" json:"nom"`
	Prenom         string           `bson:"prenom" json:"prenom"`
	Adr2           *string          `bson:"adr2,omitempty" json:"adr2,omitempty"`
	Adr4           *string          `bson:"adr4,omitempty" json:"adr4,omitempty"`
	Adr5           *string          `bson:"adr5,omitempty" json:"adr5,omitempty"`
	Pays           string           `bson:"pays" json:"pays"`
	CodePostal     string           `bson:"code_postal" json:"code_postal"`
	Ville          string           `bson:"ville" json:"ville"`
	Tel            string           `bson:"tel" json:"tel"`
	Mobile         string           `bson:"mobile" json:"mobile"`
	Email          string           `bson:"email" json:"email"`
	CodeCE         string           `bson:"Code_CE" json:"Code_CE"`
	OptinEmail     *bool            `bson:"optin_email,omitempty" json:"optin_email,omitempty"`
	OptinFidEmail  *bool            `bson:"optin_fid_email,omitempty" json:"optin_fid_email,omitempty"`
	VehiculeSimple []VehiculeSimple `bson:"vehicule_Simple,omitempty" json:"vehicule_Simple,omitempty"`
}
