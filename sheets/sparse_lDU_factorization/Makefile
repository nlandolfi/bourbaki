# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sparse_lDU_factorization.pdf

sparse_lDU_factorization.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sparse_lDU_factorization.tex

sparse_lDU_factorization.pdf: ../../*.tex ../../trademark.pdf *.tex sparse_lDU_factorization.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sparse_lDU_factorization.tex
	../../latexrun sparse_lDU_factorization.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sparse_lDU_factorization.pdf
	open sparse_lDU_factorization.pdf

o:
	make open

terms: sparse_lDU_factorization.tex
	bbk_sheet -mode ts sparse_lDU_factorization.tex

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
	rm -f sparse_lDU_factorization.*
	make reset
	bbk_sheet -mode mk > Makefile