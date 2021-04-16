# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make uniform_densities.pdf
	make graph.pdf

uniform_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > uniform_densities.tex

uniform_densities.pdf: ../../*.tex ../../trademark.pdf *.tex uniform_densities.tex
	# pdflatex --file-line-error -interaction=nonstopmode uniform_densities.tex
	../../latexrun uniform_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: uniform_densities.pdf
	open uniform_densities.pdf

o:
	make open

terms: uniform_densities.tex
	bbk_sheet -mode ts uniform_densities.tex

spell:
	aspell -c uniform_densities.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f uniform_densities.*
	make reset
	bbk_sheet -mode mk > Makefile