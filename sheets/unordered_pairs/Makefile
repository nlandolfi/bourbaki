# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make unordered_pairs.pdf

unordered_pairs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > unordered_pairs.tex

unordered_pairs.pdf: ../../*.tex ../../trademark.pdf *.tex unordered_pairs.tex
	# pdflatex --file-line-error -interaction=nonstopmode unordered_pairs.tex
	../../latexrun unordered_pairs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: unordered_pairs.pdf
	open unordered_pairs.pdf

o:
	make open

terms: unordered_pairs.tex
	bbk_sheet -mode ts unordered_pairs.tex

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
	rm -f unordered_pairs.*
	make reset
	bbk_sheet -mode mk > Makefile