# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make lDU_factorizations.pdf

lDU_factorizations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > lDU_factorizations.tex

lDU_factorizations.pdf: ../../*.tex ../../trademark.pdf *.tex lDU_factorizations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode lDU_factorizations.tex
	../../latexrun lDU_factorizations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: lDU_factorizations.pdf
	open lDU_factorizations.pdf

o:
	make open

terms: lDU_factorizations.tex
	bbk_sheet -mode ts lDU_factorizations.tex

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
	rm -f lDU_factorizations.*
	make reset
	bbk_sheet -mode mk > Makefile