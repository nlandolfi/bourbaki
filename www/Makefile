www: ../*/*.pdf ../graph.pdf
	cd .. && make graph.pdf
	cp ../graph.pdf ./
	cd ../sets && make sets.pdf
	cp ../sets/sets.pdf ./sheets/
	cd ../ordered_pairs/ && make ordered_pairs.pdf
	cp ../ordered_pairs/ordered_pairs.pdf ./sheets/
	cd ../relations/ && make relations.pdf
	cp ../relations/relations.pdf ./sheets/
	cd ../equivalence_relations/ && make equivalence_relations.pdf
	cp ../equivalence_relations/equivalence_relations.pdf ./sheets/

push: www
	cp -r * /afs/ir.stanford.edu/users/l/a/lando/WWW/bourbaki/

open: www
	make push
	open https://stanford.edu/~lando/bourbaki/
