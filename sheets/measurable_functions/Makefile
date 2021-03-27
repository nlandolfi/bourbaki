# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make measurable_functions.pdf
	make graph.pdf

measurable_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > measurable_functions.tex

measurable_functions.pdf: ../../*.tex ../../trademark.pdf *.tex measurable_functions.tex
	pdflatex --file-line-error -interaction=nonstopmode measurable_functions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measurable_functions.pdf
	open measurable_functions.pdf

o:
	make open

defs: measurable_functions.tex
	defs measurable_functions.tex

spell:
	aspell -c measurable_functions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f measurable_functions.*
	make reset
	bbk_sheet -mode mk > Makefile