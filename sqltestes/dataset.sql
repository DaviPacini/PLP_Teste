-- Inserir HEROIS no banco de dados

INSERT INTO Herois (nome_heroi, nome_real, sexo, altura, local_nasc, data_nasc, peso)
VALUES
('Homelander', 'John', 'Masculino', 1.91, 'Estados Unidos', '1982-06-10', 90.0),
('Starlight', 'Annie January', 'Feminino', 1.65, 'Estados Unidos', '1991-05-01', 55.0),
('Queen Maeve', 'Maeve', 'Feminino', 1.75, 'Estados Unidos', '1980-04-15', 70.0),
('A-Train', 'Reggie Franklin', 'Masculino', 1.77, 'Estados Unidos', '1986-03-01', 80.0),
('The Deep', 'Kevin Moskowitz', 'Masculino', 1.80, 'Estados Unidos', '1986-07-25', 85.0),
('Black Noir', 'Desconhecido', 'Masculino', 1.88, 'Desconhecido', 'Desconhecido', 100.0),
('The Soldier Boy', 'Ben', 'Masculino', 1.85, 'Estados Unidos', '1940-12-01', 95.0),
('Kimiko', 'Kimiko', 'Feminino', 1.65, 'Japão', '1985-08-10', 50.0),
('Mothers Milk', 'Marvin T. Milk', 'Masculino', 1.80, 'Estados Unidos', '1983-01-15', 90.0),
('Frenchie', 'Serge', 'Masculino', 1.75, 'França', '1980-06-20', 75.0);

-- Inserir PODERES no banco de dados

INSERT INTO Poderes (poder, descricao)
VALUES
('Voo', 'Capacidade de voar em alta velocidade e grandes distâncias.'),
('Superforça', 'Habilidade de exercer força física extraordinária, levantando grandes pesos e causando danos imensos.'),
('Visão laser', 'Capacidade de emitir feixes de laser pelos olhos com grande precisão e poder de destruição.'),
('Telecinese', 'Poder de mover objetos com a mente, sem necessidade de toque físico.'),
('Regeneração acelerada', 'Capacidade de curar ferimentos rapidamente, até mesmo regenerar membros perdidos.'),
('Supervelocidade', 'Habilidade de se mover a velocidades superiores à de um ser humano comum, com reflexos também muito rápidos.'),
('Invisibilidade', 'Capacidade de se tornar invisível aos olhos de outros.'),
('Controle de eletricidade', 'Poder de gerar e controlar eletricidade, podendo usá-la como uma arma ou para manipular dispositivos eletrônicos.'),
('Força telepática', 'Habilidade de ler mentes e comunicar-se telepaticamente com outras pessoas.'),
('Manipulação de realidade', 'Poder de alterar a realidade e os eventos que ocorrem ao redor do indivíduo, mudando as leis físicas e naturais.');

-- Inserir CRIMES no banco de dados

INSERT INTO Crimes (nome_crime, severidade)
VALUES
('Assassinato', 10),
('Roubo', 7),
('Fraude', 6),
('Sequestro', 9),
('Corrupção', 8),
('Tráfico de drogas', 9),
('Agressão física', 6),
('Estupro', 10),
('Vandalismo', 5),
('Extorsão', 8),
('Hackerismo', 7),
('Assédio sexual', 8),
('Terrorismo', 10),
('Tráfico de seres humanos', 10),
('Falsificação de documentos', 6),
('Lavagem de dinheiro', 9),
('Espionagem', 8),
('Concussão', 5),
('Fraude fiscal', 7),
('Roubos à mão armada', 9),
('Urinar em local público', 3),
('Pedofilia', 8),
('11-09', 10),
('Homofobia', 1),
('Racismo', 7),
('Trafico de animais', 5);

INSERT INTO Missoes (nome_missao, descricao, nivel_dificuldade)
VALUES
('Caçada ao Supers', 'Investigar e capturar um super que age fora da lei', 7),
('Missão no Submundo', 'Infiltrar uma rede criminosa ligada a supers', 8),
('Sabotagem Corporativa', 'Descobrir e expor segredos da Vought', 6),
('Resgate em Perigo', 'Salvar civis de um ataque descontrolado', 5),
('Operação Nocturna', 'Vigiar um super suspeito durante a noite', 3),
('Confronto Público', 'Confrontar um super em um evento televisivo', 9),
('Negociações Perigosas', 'Medir forças diplomáticas com a Vought', 7),
('Neutralizar Supers', 'Desarmar um super sem causar mortes', 8),
('Proteger Testemunha', 'Escoltar um ex-super disposto a testemunhar', 4),
('Hackeamento Crítico', 'Roubar informações sigilosas da Vought', 6),
('Reconhecimento Urbano', 'Mapear atividades suspeitas de supers na cidade', 2),
('Monitoramento Secreto', 'Acompanhar um super sem ser notado', 3),
('Entrega Segura', 'Transportar um pacote crítico sem chamar atenção', 4),
('Contato Inicial', 'Estabelecer comunicação com uma possível testemunha', 1),
('Investigação Local', 'Coletar evidências em uma cena suspeita', 3),
('Operação Silenciosa', 'Entrar e sair de um prédio sem ser detectado', 5),
('Reunião Clandestina', 'Participar de um encontro secreto sem atrair suspeitas', 4),
('Resgate Rápido', 'Retirar civis de uma área de risco antes da chegada de um super', 2),
('Alerta Comunitário', 'Informar moradores sobre a presença de um super perigoso', 1),
('Teste de Equipamento', 'Avaliar novos dispositivos contra supers', 3);
