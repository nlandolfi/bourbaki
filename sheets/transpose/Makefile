# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make transpose.pdf
	make graph.pdf

transpose.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > transpose.tex

transpose.pdf: ../../*.tex ../../trademark.pdf *.tex transpose.tex
	# pdflatex --file-line-error -interaction=nonstopmode transpose.tex
	../../latexrun transpose.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: transpose.pdf
	open transpose.pdf

o:
	make open

terms: transpose.tex
	bbk_sheet -mode ts transpose.tex

spell:
	aspell -c transpose.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f transpose.*
	make reset
	bbk_sheet -mode mk > Makefile