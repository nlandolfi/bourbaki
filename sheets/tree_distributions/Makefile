# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make tree_distributions.pdf

tree_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree_distributions.tex

tree_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex tree_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode tree_distributions.tex
	../../latexrun tree_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_distributions.pdf
	open tree_distributions.pdf

o:
	make open

terms: tree_distributions.tex
	bbk_sheet -mode ts tree_distributions.tex

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
	rm -f tree_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile