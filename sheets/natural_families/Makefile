# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_families.pdf
	make graph.pdf

natural_families.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_families.tex

natural_families.pdf: ../../*.tex ../../trademark.pdf *.tex natural_families.tex
	# pdflatex --file-line-error -interaction=nonstopmode natural_families.tex
	../../latexrun natural_families.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_families.pdf
	open natural_families.pdf

o:
	make open

terms: natural_families.tex
	bbk_sheet -mode ts natural_families.tex

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
	rm -f natural_families.*
	make reset
	bbk_sheet -mode mk > Makefile