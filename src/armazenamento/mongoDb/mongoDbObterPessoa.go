package mongoDb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testentopus/src/core/corePessoas"
	"time"
)

func (m mongoDb) Obter(id string) (*corePessoas.Pessoa, error) {
	var dbPessoa pessoa

	// converte o id de string para objectId
	objectId, errObject := primitive.ObjectIDFromHex(id)
	if errObject != nil {
		return nil, errObject
	}

	//cria um contexto para usar com o mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Ver se tem que usar isso, setar timeout ?
	defer cancel()

	// Faz a busca
	retorno := m.col.FindOne(ctx, bson.M{"_id": objectId})
	// Da o decode para o model de dbPessoa
	if err := retorno.Decode(&dbPessoa); err != nil {
		return nil, err
	}

	return pessoaMongoParaCore(&dbPessoa), nil
}
