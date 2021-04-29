# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make equivalent_sets.pdf
	make graph.pdf

equivalent_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > equivalent_sets.tex

equivalent_sets.pdf: ../../*.tex ../../trademark.pdf *.tex equivalent_sets.tex
	# pdflatex --file-line-error -interaction=nonstopmode equivalent_sets.tex
	../../latexrun equivalent_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: equivalent_sets.pdf
	open equivalent_sets.pdf

o:
	make open

terms: equivalent_sets.tex
	bbk_sheet -mode ts equivalent_sets.tex

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
	rm -f equivalent_sets.*
	make reset
	bbk_sheet -mode mk > Makefile