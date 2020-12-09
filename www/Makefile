www: ../*/*.pdf ../graph.pdf
	cd .. && make graph.pdf
	cp ../graph.pdf ./
	cd ../objects && make objects.pdf
	cp ../objects/objects.pdf ./sheets/
	cd ../sets && make sets.pdf
	cp ../sets/sets.pdf ./sheets/
	cd ../set_examples && make set_examples.pdf
	cp ../set_examples/set_examples.pdf ./sheets/
	cd ../subsets && make subsets.pdf
	cp ../subsets/subsets.pdf ./sheets/
	cd ../set_builder_notation && make set_builder_notation.pdf
	cp ../set_builder_notation/set_builder_notation.pdf ./sheets/
	cd ../ordered_pairs/ && make ordered_pairs.pdf
	cp ../ordered_pairs/ordered_pairs.pdf ./sheets/
	cd ../relations/ && make relations.pdf
	cp ../relations/relations.pdf ./sheets/
	cd ../relation_properties/ && make relation_properties.pdf
	cp ../relation_properties/relation_properties.pdf ./sheets/
	cd ../functions/ && make functions.pdf
	cp ../functions/functions.pdf ./sheets/
	cd ../function_properties/ && make function_properties.pdf
	cp ../function_properties/function_properties.pdf ./sheets/
	cd ../zero/ && make zero.pdf
	cp ../zero/zero.pdf ./sheets/

push: www
	cp -r * /afs/ir.stanford.edu/users/l/a/lando/WWW/bourbaki/

open: www
	make push
	open https://stanford.edu/~lando/bourbaki/

o:
	open index.html
