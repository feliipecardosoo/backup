package model

// Endereco representa os dados de endereço de um membro,
// contendo informações como CEP, rua, número, bairro e complemento.
//
// O campo Complemento é opcional e pode estar ausente na estrutura BSON.
type Endereco struct {
	Cep         string `bson:"cep"`                   // Código postal
	Rua         string `bson:"rua"`                   // Nome da rua
	Numero      string `bson:"numero"`                // Número da residência
	Bairro      string `bson:"bairro"`                // Nome do bairro
	Complemento string `bson:"complemento,omitempty"` // Complemento do endereço (opcional)
}

// Membro representa as informações pessoais e de status de um membro da igreja.
//
// Os campos possuem tags BSON para mapear corretamente ao banco MongoDB.
//
// Alguns campos são opcionais (como DataCasamento e NomeConjuge) e podem estar ausentes.
type Membro struct {
	Name            string   `bson:"name"`                    // Nome completo do membro
	DataNascimento  string   `bson:"dataNascimento"`          // Data de nascimento (formato string)
	AnoBatismo      int      `bson:"anoBatismo"`              // Ano em que foi batizado
	Sexo            string   `bson:"sexo"`                    // Sexo do membro
	EstadoCivil     string   `bson:"estadoCivil"`             // Estado civil atual
	DataCasamento   string   `bson:"dataCasamento,omitempty"` // Data do casamento (opcional)
	NomeConjuge     string   `bson:"nomeConjuge,omitempty"`   // Nome do cônjuge (opcional)
	Filho           bool     `bson:"filho"`                   // Indica se possui filhos
	Email           string   `bson:"email"`                   // E-mail de contato
	Telefone        string   `bson:"telefone"`                // Telefone de contato
	Status          string   `bson:"status"`                  // Status do membro (ativo, inativo, etc)
	DataStatus      string   `bson:"dataStatus"`              // Data da última alteração de status
	Validado        bool     `bson:"validado"`                // Indica se o cadastro foi validado
	Endereco        Endereco `bson:"endereco"`                // Endereço completo do membro
	DataAniversario string   `bson:"dataAniversario"`         // Data do aniversário no formato string
	DataModificacao string   `bson:"dataModificacao"`         // Data da modificacao no formato string
}
