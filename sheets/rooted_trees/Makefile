# generated automatically by sheet.go

make:
	make rooted_trees.pdf

rooted_trees.tex: sheet.tex
	sheet -mode c -in sheet.tex > rooted_trees.tex

rooted_trees.pdf: ../../*.tex ../../trademark.pdf *.tex rooted_trees.tex
	pdflatex rooted_trees.tex
	make defs


open: rooted_trees.pdf
	open rooted_trees.pdf

o:
	make open

defs: rooted_trees.tex
	defs rooted_trees.tex

spell:
	aspell -c rooted_trees.tex

s:
	make spell