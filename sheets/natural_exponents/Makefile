# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_exponents.pdf

natural_exponents.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_exponents.tex

natural_exponents.pdf: ../../*.tex ../../trademark.pdf *.tex natural_exponents.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_exponents.tex
	../../latexrun natural_exponents.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_exponents.pdf
	open natural_exponents.pdf

o:
	make open

terms: natural_exponents.tex
	bbk_sheet -mode ts natural_exponents.tex

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
	rm -f natural_exponents.*
	make reset
	bbk_sheet -mode mk > Makefile