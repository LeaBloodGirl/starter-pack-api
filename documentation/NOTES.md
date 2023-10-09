# Ressources importantes
 * La doc de l'API Tisséo : https://data.toulouse-metropole.fr/explore/dataset/api-temps-reel-tisseo/information/
 * Les configs Postman
 * La doc API Navitia :
   * https://github.com/hove-io/navitia
   * https://doc.navitia.io/#getting-started
   * https://playground.navitia.io/

 * Les principaux projets gitlab :
   * https://gitlab.tisseo.fr/ciad-iv/flask/modules-jormungandr/api-tisseo-kraken
   * https://gitlab.tisseo.fr/ciad-iv/flask/modules-jormungandr/api-pids
   * https://gitlab.tisseo.fr/ciad-iv/flask/modules-jormungandr/realtime-proxies
   * https://gitlab.tisseo.fr/ciad-iv/utilitaires/procedures/-/blob/master/navitia/tyr/admin.html

 * Les schémas :
   * https://gitlab.tisseo.fr/ciad-iv/utilitaires/documentation/-/blob/master/schemas/archi_actuelle_navitia_liens_APITiss%C3%A9o.png
   * https://gitlab.tisseo.fr/ciad-iv/utilitaires/documentation/-/blob/master/schemas/archi_fullsave_iv.png
   * https://gitlab.tisseo.fr/ciad-iv/utilitaires/documentation/-/blob/master/schemas/archi_actuelle_navitia.png
   * https://gitlab.tisseo.fr/ciad-iv/utilitaires/documentation/-/blob/master/schemas/flux_iv_temps_reel_fullsave.png
 
 * Authentification
   * https://doc.navitia.io/#about-itinerary Les fronts utilisent soit la clé dans le paramètre key, soit la clé dans le header Authentication

# Les configurations jormun/uwsgi/kraken
 * https://gitlab.tisseo.fr/ciad-iv/utilitaires/procedures/-/tree/master/navitia/jormungandr
 * Par jormungandr : sous httpd on a la conf nginx, sous uwsgi la conf uwsgi et sous configuration la conf jormun

Par exemple :
 * httpd/jormungandr_app11-5.conf => webservices.tisseo.fr sollicite la socket uwsgi jormun_5
 * uwsgi/jormungandr_app11_5.ini => uwsgi 5 charge l'appli python jormungandr_uwsgi_5, et lui autorise 5 processes maxi notamment
 * uwsgi/jormungandr_uwsgi_5.py => est le point dentrée de l'appli python jormun 5, et charge la conf jormungandr_settings_5
 * configuration/jormungandr_settings_5 => défini pas mal de constantes utilisées dans le code. On note en particuler la présence du tableau "MODULES" qui déclare les API et leur chein d'accès. Ainsi webservice.tisseo.fr/v1 pointera sur le module python 'jormungandr.modules.v1_routing.v1_routing' qui est le point d'entrée de l'API Navitia (classe présente sur gihub dans le projet Navitia). 'pids' pointe vers 'api_pids' (gitlab tisséo, projet PIDS), 'api' et 'iv' pôitent vers 'api_tisseo_kraken' (gitlab tisséo, projet api_tisseo_kraken) mais deux instances différentes, c'est à dire deux périmètres de données transport différents.

# Ressources annexes
 * https://gitlab.tisseo.fr/ciad-iv/hove-io/gonavitia
 * https://gitlab.tisseo.fr/ciad-iv/hove-io/gormungandr
 * https://github.com/govitia/navitia

# Contraintes API GO
 * Cache redis ou autre ?
 * Lire du protobuff ? Ou requêter l'API Navitia ?
 * Le format de la réponse journey change en fonction du paramètre de la requête
 * Générer un OpenAPI (swagger)
 * Gérer les logs de requêtes (gestion indépendante ou envoi dans rabbit MQ ?)
 
# Taches Macro 
 * Fil rouge : préparer l'arrivée de futurs dev (doc de prise en mains)
 * Initialiser le projet Go (archi logicielle)
 * Coder le passe-plat sans auth
 * Ecrire des tests : comparer les appels sur le Go et les appels directs sur l'API tisséo jormungandr
 * Ajouter l'auth avec tests unitaires
  * Controle des clés
  * API de gestion des utilisateurs
 * Faire le Dockefile 
 * Ecrire ensemble le process de build (avec passage Sonar + tests unitaires)
 * Déployer en Recette
 * Basculer metro_status et teleo_status
   * Donc avec lecture du protobuff ? Ou appels à API Navitia ? Ou autre ?
   * Avec gestion du cache (2 sec)

# Ressources diverses
 * http://apps.tisseo-exp.dom
 * http://apps/suivitemps/pages/login.xhtml

