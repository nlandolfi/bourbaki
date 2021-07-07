# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_differences.pdf

natural_differences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_differences.tex

natural_differences.pdf: ../../*.tex ../../trademark.pdf *.tex natural_differences.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_differences.tex
	../../latexrun natural_differences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_differences.pdf
	open natural_differences.pdf

o:
	make open

terms: natural_differences.tex
	bbk_sheet -mode ts natural_differences.tex

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
	rm -f natural_differences.*
	make reset
	bbk_sheet -mode mk > Makefile