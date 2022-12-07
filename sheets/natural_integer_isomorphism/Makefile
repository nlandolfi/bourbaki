# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_integer_isomorphism.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

natural_integer_isomorphism.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_integer_isomorphism.tex

natural_integer_isomorphism.pdf: ../../*.tex ../../trademark.pdf *.tex natural_integer_isomorphism.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_integer_isomorphism.tex
	../../latexrun natural_integer_isomorphism.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_integer_isomorphism.pdf
	open natural_integer_isomorphism.pdf

o:
	make open

terms: natural_integer_isomorphism.tex
	bbk_sheet -mode ts natural_integer_isomorphism.tex

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
	rm -f natural_integer_isomorphism.*
	make reset
	bbk_sheet -mode mk > Makefile