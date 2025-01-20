package nif

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEInformaResult(t *testing.T) {
	resp, err := GetNifDataFromEInforma(500980896)

	assert.Nil(t, err, err)

	assert.Equal(t, uint(500980896), resp.Nif)
	assert.Equal(t, uint(454850900), resp.Duns)
	assert.Equal(t, "A ARCIAL - ASSOCIAÇÃO PARA RECUPERAÇÃO DE CIDADÃOS INADAPTADOS DE OLIVEIRA DO HOSPITAL", resp.Denomination)
	assert.Equal(t, "RUA ANTÓNIO MONTEIRO", resp.Address)
	assert.Equal(t, "3400-083", resp.PostalCode)
	assert.Equal(t, "OLIVEIRA DO HOSPITAL", resp.Locality)
	assert.Equal(t, "88102 - Atividades de ação social para pessoas com incapacidades, sem alojamento", resp.Cae)
	assert.Equal(t, uint(44), resp.YearsOpen)
	assert.Equal(t, "www.arcial.pt", resp.Website)

	resp, err = GetNifDataFromEInforma(501138579)

	assert.Nil(t, err, err)

	assert.Equal(t, uint(501138579), resp.Nif)
	assert.Equal(t, uint(452889942), resp.Duns)
	assert.Equal(t, "ADAPECIL - ASSOCIAÇÃO DE AMOR PARA A EDUCAÇÃO DE CIDADÃOS INADAPTADOS DA LOURINHÃ", resp.Denomination)
	assert.Equal(t, "RUA RAINHA DONA LEONOR, 16", resp.Address)
	assert.Equal(t, "2530-922", resp.PostalCode)
	assert.Equal(t, "LOURINHÃ", resp.Locality)
	assert.Equal(t, "88102 - Atividades de ação social para pessoas com incapacidades, sem alojamento", resp.Cae)
	assert.Equal(t, uint(44), resp.YearsOpen)
	assert.Equal(t, "www.adapecil.pt", resp.Website)

	resp, err = GetNifDataFromEInforma(505014254)

	assert.Nil(t, err, err)

	assert.Equal(t, uint(505014254), resp.Nif)
	assert.Equal(t, uint(336774166), resp.Duns)
	assert.Equal(t, "CENTRO SOCIAL VALE SANTA NATÁLIA", resp.Denomination)
	assert.Equal(t, "RUA DE SÃO PEDRO, 715", resp.Address)
	assert.Equal(t, "4600-510", resp.PostalCode)
	assert.Equal(t, "ABOIM AMT", resp.Locality)
	assert.Equal(t, "87301 - Atividades de apoio social em estruturas residenciais para pessoas idosas", resp.Cae)
	assert.Equal(t, uint(24), resp.YearsOpen)
	assert.Equal(t, "www.ipss-santanatalia.com", resp.Website)
}
