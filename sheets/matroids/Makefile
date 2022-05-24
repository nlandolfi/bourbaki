# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matroids.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

matroids.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matroids.tex

matroids.pdf: ../../*.tex ../../trademark.pdf *.tex matroids.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matroids.tex
	../../latexrun matroids.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matroids.pdf
	open matroids.pdf

o:
	make open

terms: matroids.tex
	bbk_sheet -mode ts matroids.tex

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
	rm -f matroids.*
	make reset
	bbk_sheet -mode mk > Makefile