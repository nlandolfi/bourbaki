# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_distributions.pdf

product_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > product_distributions.tex

product_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex product_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_distributions.tex
	../../latexrun product_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_distributions.pdf
	open product_distributions.pdf

o:
	make open

terms: product_distributions.tex
	bbk_sheet -mode ts product_distributions.tex

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
	rm -f product_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile