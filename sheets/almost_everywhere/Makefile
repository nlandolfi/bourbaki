# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make almost_everywhere.pdf
	make graph.pdf

almost_everywhere.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > almost_everywhere.tex

almost_everywhere.pdf: ../../*.tex ../../trademark.pdf *.tex almost_everywhere.tex
	# pdflatex --file-line-error -interaction=nonstopmode almost_everywhere.tex
	../../latexrun almost_everywhere.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: almost_everywhere.pdf
	open almost_everywhere.pdf

o:
	make open

terms: almost_everywhere.tex
	bbk_sheet -mode ts almost_everywhere.tex

spell:
	aspell -c almost_everywhere.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f almost_everywhere.*
	make reset
	bbk_sheet -mode mk > Makefile