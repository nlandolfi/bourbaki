# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make line_segments.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

line_segments.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > line_segments.tex

line_segments.pdf: ../../*.tex ../../trademark.pdf *.tex line_segments.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode line_segments.tex
	../../latexrun line_segments.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: line_segments.pdf
	open line_segments.pdf

o:
	make open

terms: line_segments.tex
	bbk sheet -mode ts line_segments.tex

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
	rm -f line_segments.*
	make reset
	bbk sheet -mode mk > Makefile