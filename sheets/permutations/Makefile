# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make permutations.pdf
	make graph.pdf

permutations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > permutations.tex

permutations.pdf: ../../*.tex ../../trademark.pdf *.tex permutations.tex
	# pdflatex --file-line-error -interaction=nonstopmode permutations.tex
	../../latexrun permutations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: permutations.pdf
	open permutations.pdf

o:
	make open

terms: permutations.tex
	bbk_sheet -mode ts permutations.tex

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
	rm -f permutations.*
	make reset
	bbk_sheet -mode mk > Makefile