# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_multiplicative_identity.pdf

natural_multiplicative_identity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_multiplicative_identity.tex

natural_multiplicative_identity.pdf: ../../*.tex ../../trademark.pdf *.tex natural_multiplicative_identity.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_multiplicative_identity.tex
	../../latexrun natural_multiplicative_identity.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_multiplicative_identity.pdf
	open natural_multiplicative_identity.pdf

o:
	make open

terms: natural_multiplicative_identity.tex
	bbk_sheet -mode ts natural_multiplicative_identity.tex

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
	rm -f natural_multiplicative_identity.*
	make reset
	bbk_sheet -mode mk > Makefile