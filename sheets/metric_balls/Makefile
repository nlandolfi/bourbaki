# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make metric_balls.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

metric_balls.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > metric_balls.tex

metric_balls.pdf: ../../*.tex ../../trademark.pdf *.tex metric_balls.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode metric_balls.tex
	../../latexrun metric_balls.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_balls.pdf
	open metric_balls.pdf

o:
	make open

terms: metric_balls.tex
	bbk sheet -mode ts metric_balls.tex

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
	rm -f metric_balls.*
	make reset
	bbk sheet -mode mk > Makefile