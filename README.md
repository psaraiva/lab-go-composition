# LAB: GO COMPOSITION

## Cenário
- Uma Pessoa é alguem que nascem em algum lugar.
- Um Civil é uma Pessoa que possui um documento.
- Um Estudante é um Civil que está vinculado a uma escola.
- Um Professor é um Civil que está vinculado a uma escola e ministra uma diciplina.
- Um Militar é um Civil que tem uma patente militar.

## Interação entre os componentes
É possível acessar os dados internos de estrutura de forma direta sem fornecer o nome da estrutura interna.

Exemplo:
- Modo completo: `civilian.People.Name`
- Modo simplificado: `civilian.Name`
- Modo completo: `military.Civilian.People.Name`
- Modo simplificado: `military.Name`
- Modo completo: `military.Civilian.Document`
- Modo simplificado: `military.Document`

## Conclusão
Esse comportamento facilita o acesso a informação de forma sucinta.
