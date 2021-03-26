# Golang!

Repositório que tem como objetivo registrar os meus estudos sobre Golang.


## Como escrever código Go

Essa é a minha tradução e versão do documento [How to Write Go Code](https://golang.org/doc/code) disponível no próprio site da linguagem.

### Organização de código 

Programas em Go são organizados em **pacotes**. Um *pacote* é uma coleção de "source files" no mesmo diretório que são compilados juntos. Funções, tipos, variáveis e constantes definidas em um source file são visíveis para todos os outros source files dentro do mesmo pacote. 

Um **repositório** contém um ou mais *módulos*. Um **módulo** é uma coleção de **pacotes** Go que são lançados juntos. Um repositório Go *normalmente contém apenas um módulo*, localizado na raiz (root) do repositório. Um arquivo chamado *go.mod* declara o caminho do módulo *(module path*): O prefixo de importação para todos os pacotes dentro do módulo. O módulo contém os pacotes no diretório contendo o seu arquivo go.mod assim como os subdiretórios daquele diretório, até o próximo subdiretório contendo outro arquivo go.mod (se tiver).

Note que você não precisa publicar seu código para um repositório remoto antes de conseguir efetuar o build. **Um módulo pode ser definido localmente sem pertencer a um repositório**. Entretanto, é um bom hábito organizar seu código como se você fosse publica-lo algum dia.

Cada caminho do módulo (*module's path*) não serve apenas como um prefixo de caminho de importação para os pacotes, mas também indica onde o comando go deve olhar para efetuar o download dele. Por exemplo, para efetuar o download do módulo golang.org/x/tools, the go command would consult the repository indicated by https://golang.org/x/tools.

Um caminho de importação (*import path*) é uma string usada para importar um pacote. Um caminho de importação de um pacote é o seu caminho de importação mais seu subdiretório dentro do módulo. Por exemplo, o módulo **github.com/google/go-cmp** contém um pacote no diretório **cmp/**. O caminho de importação do pacote é **github.com/google/go-cmp/cmp**. Pacotes na biblioteca padrão não tem um prefixo do caminho de módulo.
