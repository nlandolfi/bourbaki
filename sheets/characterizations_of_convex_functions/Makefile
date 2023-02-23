# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make characterizations_of_convex_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

characterizations_of_convex_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > characterizations_of_convex_functions.tex

characterizations_of_convex_functions.pdf: ../../*.tex ../../trademark.pdf *.tex characterizations_of_convex_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode characterizations_of_convex_functions.tex
	../../latexrun characterizations_of_convex_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: characterizations_of_convex_functions.pdf
	open characterizations_of_convex_functions.pdf

o:
	make open

terms: characterizations_of_convex_functions.tex
	bbk_sheet -mode ts characterizations_of_convex_functions.tex

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
	rm -f characterizations_of_convex_functions.*
	make reset
	bbk_sheet -mode mk > Makefile