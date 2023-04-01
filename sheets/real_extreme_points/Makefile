# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_extreme_points.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_extreme_points.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_extreme_points.tex

real_extreme_points.pdf: ../../*.tex ../../trademark.pdf *.tex real_extreme_points.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_extreme_points.tex
	../../latexrun real_extreme_points.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_extreme_points.pdf
	open real_extreme_points.pdf

o:
	make open

terms: real_extreme_points.tex
	bbk_sheet -mode ts real_extreme_points.tex

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
	rm -f real_extreme_points.*
	make reset
	bbk_sheet -mode mk > Makefile