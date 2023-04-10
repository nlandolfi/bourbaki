# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make inverses_of_composite_relations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

inverses_of_composite_relations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > inverses_of_composite_relations.tex

inverses_of_composite_relations.pdf: ../../*.tex ../../trademark.pdf *.tex inverses_of_composite_relations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode inverses_of_composite_relations.tex
	../../latexrun inverses_of_composite_relations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inverses_of_composite_relations.pdf
	open inverses_of_composite_relations.pdf

o:
	make open

terms: inverses_of_composite_relations.tex
	bbk sheet -mode ts inverses_of_composite_relations.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f inverses_of_composite_relations.*
	make reset
	bbk sheet -mode mk > Makefile