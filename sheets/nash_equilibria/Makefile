# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make nash_equilibria.pdf
	make graph.pdf

nash_equilibria.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > nash_equilibria.tex

nash_equilibria.pdf: ../../*.tex ../../trademark.pdf *.tex nash_equilibria.tex
	pdflatex --file-line-error -interaction=nonstopmode nash_equilibria.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: nash_equilibria.pdf
	open nash_equilibria.pdf

o:
	make open

defs: nash_equilibria.tex
	defs nash_equilibria.tex

spell:
	aspell -c nash_equilibria.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f nash_equilibria.*
	make reset
	bbk_sheet -mode mk > Makefile