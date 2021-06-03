# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make identity_matrices.pdf

identity_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > identity_matrices.tex

identity_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex identity_matrices.tex
	# pdflatex --file-line-error -interaction=nonstopmode identity_matrices.tex
	../../latexrun identity_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: identity_matrices.pdf
	open identity_matrices.pdf

o:
	make open

terms: identity_matrices.tex
	bbk_sheet -mode ts identity_matrices.tex

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
	rm -f identity_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile