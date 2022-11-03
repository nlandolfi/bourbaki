# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make submatrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

submatrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > submatrices.tex

submatrices.pdf: ../../*.tex ../../trademark.pdf *.tex submatrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode submatrices.tex
	../../latexrun submatrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: submatrices.pdf
	open submatrices.pdf

o:
	make open

terms: submatrices.tex
	bbk_sheet -mode ts submatrices.tex

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
	rm -f submatrices.*
	make reset
	bbk_sheet -mode mk > Makefile