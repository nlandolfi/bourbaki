# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make product_measures.pdf
	make graph.pdf

product_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > product_measures.tex

product_measures.pdf: ../../*.tex ../../trademark.pdf *.tex product_measures.tex
	# pdflatex --file-line-error -interaction=nonstopmode product_measures.tex
	../../latexrun product_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_measures.pdf
	open product_measures.pdf

o:
	make open

terms: product_measures.tex
	bbk_sheet -mode ts product_measures.tex

spell:
	aspell -c product_measures.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f product_measures.*
	make reset
	bbk_sheet -mode mk > Makefile