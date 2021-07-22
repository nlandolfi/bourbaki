# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make counting_measures.pdf

counting_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > counting_measures.tex

counting_measures.pdf: ../../*.tex ../../trademark.pdf *.tex counting_measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode counting_measures.tex
	../../latexrun counting_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: counting_measures.pdf
	open counting_measures.pdf

o:
	make open

terms: counting_measures.tex
	bbk_sheet -mode ts counting_measures.tex

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
	rm -f counting_measures.*
	make reset
	bbk_sheet -mode mk > Makefile