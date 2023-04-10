# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make alphabets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

alphabets.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > alphabets.tex

alphabets.pdf: ../../*.tex ../../trademark.pdf *.tex alphabets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode alphabets.tex
	../../latexrun alphabets.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: alphabets.pdf
	open alphabets.pdf

o:
	make open

terms: alphabets.tex
	bbk sheet -mode ts alphabets.tex

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
	rm -f alphabets.*
	make reset
	bbk sheet -mode mk > Makefile