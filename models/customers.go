package models

type Vehicule struct {
	LibelleCourtGamme *string `bson:"libelle_court_gamme,omitempty" json:"libelle_court_gamme,omitempty"`
}

type VehiculeSimple struct {
	CodeVin           *string `bson:"code_vin,omitempty" json:"code_vin,omitempty"`
	LibelleGamme      *string `bson:"libelle_gamme,omitempty" json:"libelle_gamme,omitempty"`
	LibelleCourtGamme *string `bson:"libelle_court_gamme,omitempty" json:"libelle_court_gamme,omitempty"`
	Carrosserie       *string `bson:"carrosserie,omitempty" json:"carrosserie,omitempty"`
	Energie           *string `bson:"energierie,omitempty" json:"energierie,omitempty"`
	LibelleCouleur    *string `bson:"libelle_couleur,omitempty" json:"libelle_couleur,omitempty"`
	DateAchatVehicule *string `bson:"date_achat_vehicule,omitempty" json:"date_achat_vehicule,omitempty"`
	DatePremiereImmat *string `bson:"date_premiere_immat,omitempty" json:"date_premiere_immat,omitempty"`
	Kilometrage       *string `bson:"kilometrage,omitempty" json:"kilometrage,omitempty"`
	DateEntretien     *string `bson:"date_entretien,omitempty" json:"date_entretien,omitempty"`
	DateVente         *string `bson:"date_vente,omitempty" json:"date_vente,omitempty"`
	TypeAchat         *string `bson:"type_achat,omitempty" json:"type_achat,omitempty"`
}

type Customers struct {
	IDClient       uint32     `bson:"id_client" json:"id_client"`
	Nom            string     `bson:"nom" json:"nom"`
	Prenom         string     `bson:"prenom" json:"prenom"`
	Adr2           *string    `bson:"adr2,omitempty" json:"adr2,omitempty"`
	Adr4           *string    `bson:"adr4,omitempty" json:"adr4,omitempty"`
	Adr5           *string    `bson:"adr5,omitempty" json:"adr5,omitempty"`
	Pays           string     `bson:"pays" json:"pays"`
	CodePostal     string     `bson:"code_postal" json:"code_postal"`
	Ville          string     `bson:"ville" json:"ville"`
	Tel            string     `bson:"tel" json:"tel"`
	Mobile         string     `bson:"mobile" json:"mobile"`
	Email          string     `bson:"email" json:"email"`
	CodeCE         string     `bson:"Code_CE" json:"Code_CE"`
	OptinEmail     *bool      `bson:"optin_email,omitempty" json:"optin_email,omitempty"`
	OptinFidEmail  *bool      `bson:"optin_fid_email,omitempty" json:"optin_fid_email,omitempty"`
	VehiculeSimple []Vehicule `bson:"vehicule_Simple,omitempty" json:"vehicule_Simple,omitempty"`
}

type Customer struct {
	IDClient           uint32            `bson:"id_client" json:"id_client"`
	Nom                string            `bson:"nom" json:"nom"`
	Prenom             string            `bson:"prenom" json:"prenom"`
	Adr2               *string           `bson:"adr2,omitempty" json:"adr2,omitempty"`
	Adr4               *string           `bson:"adr4,omitempty" json:"adr4,omitempty"`
	Adr5               *string           `bson:"adr5,omitempty" json:"adr5,omitempty"`
	Pays               string            `bson:"pays" json:"pays"`
	CodePostal         string            `bson:"code_postal" json:"code_postal"`
	Ville              string            `bson:"ville" json:"ville"`
	Tel                string            `bson:"tel" json:"tel"`
	Mobile             string            `bson:"mobile" json:"mobile"`
	Email              string            `bson:"email" json:"email"`
	CodeCE             string            `bson:"Code_CE" json:"Code_CE"`
	CodeCEDecla        *string           `bson:"code_ce_decla,omitempty" json:"code_ce_decla,omitempty"`
	OptinEmail         *bool             `bson:"optin_email,omitempty" json:"optin_email,omitempty"`
	OptinFidEmail      *bool             `bson:"optin_fid_email,omitempty" json:"optin_fid_email,omitempty"`
	DateNaissance      *string           `bson:"date_naissance,omitempty" json:"date_naissance,omitempty"`
	MomentJoindre      *string           `bson:"Moment_joindre,omitempty" json:"Moment_joindre,omitempty"`
	DemarchageTel      *bool             `bson:"Demarchage_Tel,omitempty" json:"Demarchage_Tel,omitempty"`
	DemarchageMobile   *bool             `bson:"Demarchage_mobile,omitempty" json:"Demarchage_mobile,omitempty"`
	Statut             *string           `bson:"statut,omitempty" json:"statut,omitempty"`
	OptinSMS           *bool             `bson:"optin_SMS,omitempty" json:"optin_SMS,omitempty"`
	BloctelTel         *bool             `bson:"Bloctel_tel,omitempty" json:"Bloctel_tel,omitempty"`
	BloctelMobile      *bool             `bson:"Bloctel_mobile,omitempty" json:"Bloctel_mobile,omitempty"`
	SourceOrigine      *string           `bson:"Source_origine,omitempty" json:"Source_origine,omitempty"`
	MentionLegale      *string           `bson:"Mention_legale,omitempty" json:"Mention_legale,omitempty"`
	Inscription        *string           `bson:"inscription,omitempty" json:"inscription,omitempty"`
	PreInscription     *string           `bson:"pre_inscription,omitempty" json:"pre_inscription,omitempty"`
	VehiculeConcurrent *interface{}      `bson:"vehicule_concurrent,omitempty" json:"vehicule_concurrent,omitempty"`
	ContactsEntreprise *interface{}      `bson:"contacts_entreprise,omitempty" json:"contacts_entreprise,omitempty"`
	VehiculeSimple     *[]VehiculeSimple `bson:"vehicule_Simple,omitempty" json:"vehicule_Simple,omitempty"`
}
