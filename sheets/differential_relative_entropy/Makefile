# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make differential_relative_entropy.pdf
	make graph.pdf

differential_relative_entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > differential_relative_entropy.tex

differential_relative_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_relative_entropy.tex
	pdflatex --file-line-error -interaction=nonstopmode differential_relative_entropy.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_relative_entropy.pdf
	open differential_relative_entropy.pdf

o:
	make open

defs: differential_relative_entropy.tex
	defs differential_relative_entropy.tex

spell:
	aspell -c differential_relative_entropy.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f differential_relative_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile