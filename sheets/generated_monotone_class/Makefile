# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make generated_monotone_class.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

generated_monotone_class.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > generated_monotone_class.tex

generated_monotone_class.pdf: ../../*.tex ../../trademark.pdf *.tex generated_monotone_class.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode generated_monotone_class.tex
	../../latexrun generated_monotone_class.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generated_monotone_class.pdf
	open generated_monotone_class.pdf

o:
	make open

terms: generated_monotone_class.tex
	bbk sheet -mode ts generated_monotone_class.tex

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
	rm -f generated_monotone_class.*
	make reset
	bbk sheet -mode mk > Makefile