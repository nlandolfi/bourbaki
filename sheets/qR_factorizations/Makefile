# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make qR_factorizations.pdf

qR_factorizations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > qR_factorizations.tex

qR_factorizations.pdf: ../../*.tex ../../trademark.pdf *.tex qR_factorizations.tex
	# pdflatex --file-line-error -interaction=nonstopmode qR_factorizations.tex
	../../latexrun qR_factorizations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: qR_factorizations.pdf
	open qR_factorizations.pdf

o:
	make open

terms: qR_factorizations.tex
	bbk_sheet -mode ts qR_factorizations.tex

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
	rm -f qR_factorizations.*
	make reset
	bbk_sheet -mode mk > Makefile