# Go Lang

[![NPM](https://img.shields.io/npm/l/react)](https://github.com/wesloy/Portifolio_S.O.L.I.D/blob/main/license)

## Sobre o projeto

Este projeto tem por base desvendar e aplicar conceitos da linguagem _Go Lang_.  
Irei conceituar características do _Go_ e aplicar durante a construção de uma API de controles de leitura de documentos.

**Observação**: toda palavra seguida de asterisco, possui uma breve descrição na penúltima sessão deste _Redme_, sessão _Glossário_, com o objetivo de elucidar quaisquer dúvidas sobre a utilização do termo no contexto deste artigo.  
Exemplo: Glossário\*

## Go Lang

Uma linguagem desenvolvida pelo Google, open source* e tem como objetivo tornar o programador mais eficiente, rápido e produtivo.  
Linguagem limpa, baixa curva de apredizagem, madura e simples. A Go já foi criada com foco multicore* e multiplataforma* ela é estaticamente tipada*, com baixa verbosidade* e precisa ser compilada*, mas o resultado é um único arquivo.

### Instalação

https://go.dev/ sendo este o site oficial, pode-se baixar o instalador para qualquer das 3 plataformas comuns de sistemas operacionais: Linux, Mac ou Windows.  
A instalação segue aquele "burocrático" procedimento: Next, next e finish. Easy-peasy!

### Hello World

O "go.mod" é um arquivo gerado com o comando destacado abaixo e este arquivo gerencia todas as dependências importadas para o projeto, bem como as dependências que poderão ser criadas no projeto em que você estiver "codando".  
Atualmente o "go.mod" também contem a localização/caminho do repositório do código fonte que está sendo desenvolvido, podendo ser a raiz do servidor onde será feito o deploy ou um repositório do github, que aliás é o que mais se é feito, principalmente se a intenção é publicar para que outros possam usufruir da aplicação criada.

Após criar a pasta onde o projeto será desenvolvido e acessar a mesma através do seu prompt de comando, digite a seguinte instrução:

```DOS
go mod init example/hello
```

O comando "go mod init" é do Go Lang, já o sufixo "example/hello" é o que se foi explicado acima, ou seja, o endereço principal que sua aplicação adotará, que muitas vezes é o nome do repositório git.  

![go mod init](imgs/go_mod_init.png)

Após a criação e inicialização do projeto, precisamos *codar* o esperado ___Hello World___.  
Para isso, precisamos criar um arquivo com extensão *.go e inserir as instruções necessárias para exibir a mensagem.  

![sintaxe hello world](imgs/hello_world_sintaxe.png)  

A sintaxe da imagem é esta:  

```go
package main 

import "fmt" /* Importado de uma biblioteca GO */

func main() { 
	fmt.Println("Hello, World!") 
}
```

Até mesmo como a imagem mostra, nos comentários inseridos no código, temos alguns elementos importantíssimos para um _Dev Go_. São eles:  

- ***package main***: ou pacote principal/essencial. Todo "package" é um aglotinador de métodos, arquivos e funções. Neste caso, em específico, o pacote principal do projeto.  
- ***import fmt***: trata-se de uma das várias bibliotecas que o GO fornece em sua instalação de forma padrão. O _fmt_ é uma biblioteca que fornece formatações, até mesmo a possibilidade de imprimir em tela um texto, método que usamos em nosso exemplo.  
- ***func main***: a função principal/fundamental, como em outras linguagens (C# ou Java) é necessária para iniciar a execução do programa, sem ela não é possível se rodar a aplicação.  

Agora é fazer a mágica acontecer!  
Vamos executar nosso projeto e dar um "Olá mundo!".  
Para isso, vamos abrir o terminal e usar o seguinte comando:  

```go
go run .
```
 Deixando o próprio Go encontrar o arquivo que contem a estrutura principal (Main).  
 Pode-se especificar o arquivo a ser executado, com o comando abaixo:  
 ```go
go run hello.go
```

![run go](imgs/run_go.png)  

### Chamar códigos de um arquivo externo/diferente  


Até aqui vimos o *Hello World*, sendo executado dentro do arquivo principal e dentro da função principal, denominada *Main*. Mas não seria razoável pensar em "codar" todo um sistema/projeto em um único arquivo, então se faz necessário entender como podemos chamar códigos de um *package* diferente, que esteja dentro de um arquivo externo ao *Main* (Principal).  

**Imagem 01**

![func in modules p1](imgs/hello_world_func.png)

Na **Imagem 01** foi dado destaque a estrutura de uma função em arquivos distintos, os passos para codar o que está nesta imagem são:  

- Primeiramente o "go mod init", que já citamos acima, deve iniciar com um endereço Git ou de pastas que você está codando. Depois montamos o arquivo principal, como no "Hello World" exemplificado no tópico anterior, que nesta **Imagem 01** foi nomeado como "main.go".
- Agora nos resta criar um arquivo à parte com a função que se pretende rodar, porém é preciso detida atenção neste momento, para entender o jeto Go de estruturar arquivos externos.
  
**Imagem 02**

![func in modules p2](imgs/hello_world_func2.png)  

- Na **Imagem 02** podemos observar que o "_package_" é formado quando criamos uma subpasta e dentro dela criamos o arquivo com extensão *.go.  Para exemplificação criamos o arquivo "hello.go" e "bye.go" eles possuem como "cabeçalho" o "package mensages" o que não é uma coincidência, mas sim o jeto Go de marcar sua estrutura de relacionamento de arquivos, veja na imagem abaixo.  

![func in modules p3](imgs/hello_world_func3.png)  

## Links Úteis / Referências

- [Documentação oficial Go Lang](https://go.dev/)  
- [Canal Youtube - Wesley Willians - Go Lang do Zero](https://www.youtube.com/watch?v=_MkQLDMak-4&list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg)  


## Glossário

- **Glossário** - dicionário de palavras de sentido obscuro ou pouco conhecido; elucidário.
- **Estaticamente Tipada** - quando a pessoa que está programando precisa informar explicitamente o tipo de cada dado utilizado no sistema.
- **Multiplataforma** - mais de 1 sistema operacional pode rodar a aplicação. Exemplo: Linux, Mac ou Windows.
- **Multicore** - contém pelo menos 2 núcleos de processamento.
- **Open Source** - código aberto, onde os usuários podem contribuir para a sua evolução e utilzar para qualquer fim.
- **Verbosidade** - necessário gerar muito código para que um processo seja executado.

## Autor

> **Wesley Eloy**  
> Bacharel em Administração  
> Pós graduado em Tecnologias de aplicações Web  
> Atuando como Desenvolvedor  
> [Linkdin (wesley-eloy)](https://www.linkedin.com/in/wesley-eloy/)




