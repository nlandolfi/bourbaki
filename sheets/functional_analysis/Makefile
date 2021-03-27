# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make functional_analysis.pdf
	make graph.pdf

functional_analysis.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > functional_analysis.tex

functional_analysis.pdf: ../../*.tex ../../trademark.pdf *.tex functional_analysis.tex
	pdflatex --file-line-error -interaction=nonstopmode functional_analysis.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: functional_analysis.pdf
	open functional_analysis.pdf

o:
	make open

defs: functional_analysis.tex
	defs functional_analysis.tex

spell:
	aspell -c functional_analysis.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f functional_analysis.*
	make reset
	bbk_sheet -mode mk > Makefile