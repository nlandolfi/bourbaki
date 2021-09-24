# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_random_functions.pdf

normal_random_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > normal_random_functions.tex

normal_random_functions.pdf: ../../*.tex ../../trademark.pdf *.tex normal_random_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normal_random_functions.tex
	../../latexrun normal_random_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_random_functions.pdf
	open normal_random_functions.pdf

o:
	make open

terms: normal_random_functions.tex
	bbk_sheet -mode ts normal_random_functions.tex

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
	rm -f normal_random_functions.*
	make reset
	bbk_sheet -mode mk > Makefile