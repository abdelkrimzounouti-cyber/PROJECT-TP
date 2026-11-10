# Documentation du projet PCG

# 1. Analyse lexicale : 
c'est la premier partie de notre compilateur et Elle permet de lire le code source et de le découper en unités appelées tokens.

Elle reconnaît notamment :
- Les mots-clés (`program`, `begin`, `end`, `if`, `while`...)
- Les identificateurs (noms des variables)
- Les nombres et les valeurs booléennes
- Les opérateurs arithmétiques et logiques
- Les symboles utilisés dans le langage.

Les espaces et les commentaires sont ignorés. Les caractères non reconnus provoquent des erreurs lexicales.

# 2. Analyse syntaxique : 

la 2eme partie c'est L'analyse syntaxique qui vérifie que le programme respecte les règles de grammaire du langage PCG.

Elle permet notamment de vérifier :
- La structure générale d'un programme.
- Les déclarations des variables.
- Les affectations.
- Les instructions conditionnelles (`if`).
- Les boucles (`while`).
- Les expressions arithmétiques et logiques.

Une erreur est signalée lorsque le programme ne respecte pas la grammaire définie dans `Pcg.g4`.

# 3. Analyse sémantique : 

la 3eme partie c'est L'analyse sémantique qui vérifie que le programme est cohérent du point de vue des types(typage).

Elle utilise une table des symboles pour enregistrer les variables déclarées et leurs types.

Elle vérifie notamment :
- Que les variables utilisées sont déclarées.
- Que les affectations respectent les types.
- Que les opérations sont compatibles avec les types des opérandes.
- Que les conditions des instructions `if` et `while` sont booléennes.

Les types principaux utilisés sont `Integer` et `Boolean`.

Une erreur sémantique est signalée lorsqu'une opération ou une affectation n'est pas compatible avec les types concernés.

## 4. Compilation et tests : 

On a utiliser Go ( plus precisement dans la partie typage ou on a modifier quelque fonctions comme typecheck_decl() , typecheck_inst() et  typecheck_expr() dans le fichier typing.go ) et aussi ANTLR (pour la parite d'analyse lexicale ) pour construire le compilateur.

Le répertoire `Tests/` contient des programmes corrects et incorrects permettant de vérifier les trois premier étapes du compilateur.

Les fichiers d'erreur suivent les conventions suivantes :
- `l_erreur.pcg` : erreur lexicale.
- `s_erreur.pcg` : erreur syntaxique.
- `t_erreur.pcg` : erreur sémantique.

Et Les commandes suivantes permettent de générer et d'afficher les arbres des programmes de test : 

Arbres syntaxiques(sans typage) :

 Test syntaxique 1 : draw-tree syntax1.dot
 Test syntaxique 2 : draw-tree syntax2.dot
 
Arbres avec informations de typage :

Test de typage 1 : draw-tree --show-attrs t typing1.dot
Test de typage 2 : draw-tree --show-attrs t typing2.dot
Test de typage 3 (erreur sémantique) : draw-tree --show-attrs t typing3.dot
