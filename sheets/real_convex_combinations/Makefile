# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_convex_combinations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_convex_combinations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_convex_combinations.tex

real_convex_combinations.pdf: ../../*.tex ../../trademark.pdf *.tex real_convex_combinations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_convex_combinations.tex
	../../latexrun real_convex_combinations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_convex_combinations.pdf
	open real_convex_combinations.pdf

o:
	make open

terms: real_convex_combinations.tex
	bbk_sheet -mode ts real_convex_combinations.tex

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
	rm -f real_convex_combinations.*
	make reset
	bbk_sheet -mode mk > Makefile