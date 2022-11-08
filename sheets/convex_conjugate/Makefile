# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convex_conjugate.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

convex_conjugate.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > convex_conjugate.tex

convex_conjugate.pdf: ../../*.tex ../../trademark.pdf *.tex convex_conjugate.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode convex_conjugate.tex
	../../latexrun convex_conjugate.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convex_conjugate.pdf
	open convex_conjugate.pdf

o:
	make open

terms: convex_conjugate.tex
	bbk_sheet -mode ts convex_conjugate.tex

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
	rm -f convex_conjugate.*
	make reset
	bbk_sheet -mode mk > Makefile