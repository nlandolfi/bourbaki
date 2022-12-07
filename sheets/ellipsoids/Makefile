# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ellipsoids.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

ellipsoids.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > ellipsoids.tex

ellipsoids.pdf: ../../*.tex ../../trademark.pdf *.tex ellipsoids.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ellipsoids.tex
	../../latexrun ellipsoids.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ellipsoids.pdf
	open ellipsoids.pdf

o:
	make open

terms: ellipsoids.tex
	bbk_sheet -mode ts ellipsoids.tex

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
	rm -f ellipsoids.*
	make reset
	bbk_sheet -mode mk > Makefile