# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make alphabets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

alphabets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > alphabets.tex

alphabets.pdf: ../../*.tex ../../trademark.pdf *.tex alphabets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode alphabets.tex
	../../latexrun alphabets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: alphabets.pdf
	open alphabets.pdf

o:
	make open

terms: alphabets.tex
	bbk_sheet -mode ts alphabets.tex

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
	bbk_sheet -mode mk > Makefile