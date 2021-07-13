# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rational_products.pdf

rational_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rational_products.tex

rational_products.pdf: ../../*.tex ../../trademark.pdf *.tex rational_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rational_products.tex
	../../latexrun rational_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_products.pdf
	open rational_products.pdf

o:
	make open

terms: rational_products.tex
	bbk_sheet -mode ts rational_products.tex

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
	rm -f rational_products.*
	make reset
	bbk_sheet -mode mk > Makefile