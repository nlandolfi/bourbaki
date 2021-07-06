# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make intersection_of_empty_set.pdf

intersection_of_empty_set.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > intersection_of_empty_set.tex

intersection_of_empty_set.pdf: ../../*.tex ../../trademark.pdf *.tex intersection_of_empty_set.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode intersection_of_empty_set.tex
	../../latexrun intersection_of_empty_set.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: intersection_of_empty_set.pdf
	open intersection_of_empty_set.pdf

o:
	make open

terms: intersection_of_empty_set.tex
	bbk_sheet -mode ts intersection_of_empty_set.tex

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
	rm -f intersection_of_empty_set.*
	make reset
	bbk_sheet -mode mk > Makefile