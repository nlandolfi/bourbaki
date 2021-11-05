# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monotonic_functions.pdf

monotonic_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monotonic_functions.tex

monotonic_functions.pdf: ../../*.tex ../../trademark.pdf *.tex monotonic_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monotonic_functions.tex
	../../latexrun monotonic_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotonic_functions.pdf
	open monotonic_functions.pdf

o:
	make open

terms: monotonic_functions.tex
	bbk_sheet -mode ts monotonic_functions.tex

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
	rm -f monotonic_functions.*
	make reset
	bbk_sheet -mode mk > Makefile