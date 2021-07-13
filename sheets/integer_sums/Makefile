# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_sums.pdf

integer_sums.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integer_sums.tex

integer_sums.pdf: ../../*.tex ../../trademark.pdf *.tex integer_sums.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_sums.tex
	../../latexrun integer_sums.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_sums.pdf
	open integer_sums.pdf

o:
	make open

terms: integer_sums.tex
	bbk_sheet -mode ts integer_sums.tex

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
	rm -f integer_sums.*
	make reset
	bbk_sheet -mode mk > Makefile