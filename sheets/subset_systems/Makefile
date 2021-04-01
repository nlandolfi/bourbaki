# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make subset_systems.pdf
	make graph.pdf

subset_systems.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subset_systems.tex

subset_systems.pdf: ../../*.tex ../../trademark.pdf *.tex subset_systems.tex
	pdflatex --file-line-error -interaction=nonstopmode subset_systems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subset_systems.pdf
	open subset_systems.pdf

o:
	make open

terms: subset_systems.tex
	bbk_sheet -mode ts subset_systems.tex

spell:
	aspell -c subset_systems.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f subset_systems.*
	make reset
	bbk_sheet -mode mk > Makefile