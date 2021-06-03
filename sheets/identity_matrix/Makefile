# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make identity_matrix.pdf

identity_matrix.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > identity_matrix.tex

identity_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex identity_matrix.tex
	# pdflatex --file-line-error -interaction=nonstopmode identity_matrix.tex
	../../latexrun identity_matrix.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: identity_matrix.pdf
	open identity_matrix.pdf

o:
	make open

terms: identity_matrix.tex
	bbk_sheet -mode ts identity_matrix.tex

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
	rm -f identity_matrix.*
	make reset
	bbk_sheet -mode mk > Makefile